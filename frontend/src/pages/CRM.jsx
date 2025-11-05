import { useState, useEffect } from 'react';
import { crmAPI } from '../services/api';
import { Building2, Phone, Mail, User, Plus, MessageSquare } from 'lucide-react';
import toast from 'react-hot-toast';

export default function CRM() {
  const [customers, setCustomers] = useState([]);
  const [selectedCustomer, setSelectedCustomer] = useState(null);
  const [interactions, setInteractions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showInteractionForm, setShowInteractionForm] = useState(false);

  useEffect(() => {
    loadCustomers();
  }, []);

  useEffect(() => {
    if (selectedCustomer) {
      loadInteractions(selectedCustomer.id);
    }
  }, [selectedCustomer]);

  const loadCustomers = async () => {
    setLoading(true);
    try {
      const response = await crmAPI.listCustomers({ page: 1, limit: 50 });
      setCustomers(response.data.data || []);
    } catch (error) {
      toast.error('Failed to load customers');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const loadInteractions = async (customerId) => {
    try {
      const response = await crmAPI.listInteractionsByCustomer(customerId, {
        page: 1,
        limit: 50,
      });
      setInteractions(response.data.data || []);
    } catch (error) {
      toast.error('Failed to load interactions');
      console.error(error);
    }
  };

  const handleAddInteraction = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = {
      customer_id: selectedCustomer.id,
      interaction_type: formData.get('type'),
      subject: formData.get('subject'),
      notes: formData.get('notes'),
    };

    try {
      await crmAPI.createInteraction(data);
      toast.success('Interaction added successfully');
      setShowInteractionForm(false);
      loadInteractions(selectedCustomer.id);
      e.target.reset();
    } catch (error) {
      toast.error('Failed to add interaction');
      console.error(error);
    }
  };

  const getStatusColor = (status) => {
    const colors = {
      lead: 'bg-yellow-100 text-yellow-800',
      prospect: 'bg-blue-100 text-blue-800',
      active: 'bg-green-100 text-green-800',
      inactive: 'bg-gray-100 text-gray-800',
    };
    return colors[status] || 'bg-gray-100 text-gray-800';
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">CRM</h1>
          <p className="text-gray-600 mt-1">Manage customer relationships and interactions</p>
        </div>
      </div>

      {loading ? (
        <div className="text-center py-12">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600 mx-auto"></div>
          <p className="text-gray-500 mt-4">Loading...</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Customers List */}
          <div className="lg:col-span-1">
            <div className="bg-white rounded-lg shadow">
              <div className="px-6 py-4 border-b border-gray-200">
                <h2 className="text-lg font-semibold">Customers</h2>
              </div>
              <div className="divide-y divide-gray-200 max-h-[600px] overflow-y-auto">
                {customers.length === 0 ? (
                  <div className="px-6 py-12 text-center text-gray-500">
                    No customers found
                  </div>
                ) : (
                  customers.map((customer) => (
                    <div
                      key={customer.id}
                      onClick={() => setSelectedCustomer(customer)}
                      className={`px-6 py-4 cursor-pointer hover:bg-gray-50 ${
                        selectedCustomer?.id === customer.id ? 'bg-indigo-50' : ''
                      }`}
                    >
                      <div className="flex items-start justify-between">
                        <div className="flex-1">
                          <div className="flex items-center">
                            <Building2 className="w-5 h-5 text-gray-400 mr-2" />
                            <h3 className="text-sm font-medium text-gray-900">
                              {customer.company_name}
                            </h3>
                          </div>
                          {customer.contact_name && (
                            <p className="text-sm text-gray-500 mt-1 ml-7">
                              {customer.contact_name}
                            </p>
                          )}
                          <div className="mt-2">
                            <span className={`px-2 py-1 text-xs rounded-full ${getStatusColor(customer.status)}`}>
                              {customer.status}
                            </span>
                          </div>
                        </div>
                      </div>
                    </div>
                  ))
                )}
              </div>
            </div>
          </div>

          {/* Customer Details & Interactions */}
          <div className="lg:col-span-2">
            {selectedCustomer ? (
              <div className="space-y-6">
                {/* Customer Details */}
                <div className="bg-white rounded-lg shadow p-6">
                  <h2 className="text-xl font-semibold mb-4">{selectedCustomer.company_name}</h2>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {selectedCustomer.contact_name && (
                      <div className="flex items-center">
                        <User className="w-5 h-5 text-gray-400 mr-2" />
                        <span className="text-gray-600">{selectedCustomer.contact_name}</span>
                      </div>
                    )}
                    {selectedCustomer.email && (
                      <div className="flex items-center">
                        <Mail className="w-5 h-5 text-gray-400 mr-2" />
                        <span className="text-gray-600">{selectedCustomer.email}</span>
                      </div>
                    )}
                    {selectedCustomer.phone && (
                      <div className="flex items-center">
                        <Phone className="w-5 h-5 text-gray-400 mr-2" />
                        <span className="text-gray-600">{selectedCustomer.phone}</span>
                      </div>
                    )}
                    {selectedCustomer.industry && (
                      <div className="flex items-center">
                        <Building2 className="w-5 h-5 text-gray-400 mr-2" />
                        <span className="text-gray-600">{selectedCustomer.industry}</span>
                      </div>
                    )}
                  </div>
                </div>

                {/* Interactions */}
                <div className="bg-white rounded-lg shadow">
                  <div className="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
                    <h2 className="text-lg font-semibold">Interactions</h2>
                    <button
                      onClick={() => setShowInteractionForm(!showInteractionForm)}
                      className="flex items-center px-3 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700"
                    >
                      <Plus className="w-4 h-4 mr-1" />
                      Add Interaction
                    </button>
                  </div>

                  {showInteractionForm && (
                    <div className="px-6 py-4 bg-gray-50 border-b border-gray-200">
                      <form onSubmit={handleAddInteraction} className="space-y-4">
                        <div>
                          <label className="block text-sm font-medium text-gray-700">Type</label>
                          <select
                            name="type"
                            required
                            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                          >
                            <option value="call">Call</option>
                            <option value="email">Email</option>
                            <option value="meeting">Meeting</option>
                            <option value="note">Note</option>
                          </select>
                        </div>
                        <div>
                          <label className="block text-sm font-medium text-gray-700">Subject</label>
                          <input
                            type="text"
                            name="subject"
                            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                          />
                        </div>
                        <div>
                          <label className="block text-sm font-medium text-gray-700">Notes</label>
                          <textarea
                            name="notes"
                            rows={3}
                            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                          ></textarea>
                        </div>
                        <div className="flex space-x-2">
                          <button
                            type="submit"
                            className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700"
                          >
                            Save
                          </button>
                          <button
                            type="button"
                            onClick={() => setShowInteractionForm(false)}
                            className="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300"
                          >
                            Cancel
                          </button>
                        </div>
                      </form>
                    </div>
                  )}

                  <div className="divide-y divide-gray-200 max-h-[400px] overflow-y-auto">
                    {interactions.length === 0 ? (
                      <div className="px-6 py-12 text-center text-gray-500">
                        No interactions yet
                      </div>
                    ) : (
                      interactions.map((interaction) => (
                        <div key={interaction.id} className="px-6 py-4">
                          <div className="flex items-start">
                            <MessageSquare className="w-5 h-5 text-gray-400 mr-3 mt-1" />
                            <div className="flex-1">
                              <div className="flex items-center justify-between">
                                <span className="text-xs font-medium text-gray-500 uppercase">
                                  {interaction.interaction_type}
                                </span>
                                <span className="text-xs text-gray-500">
                                  {new Date(interaction.interaction_date).toLocaleDateString()}
                                </span>
                              </div>
                              {interaction.subject && (
                                <h4 className="text-sm font-medium text-gray-900 mt-1">
                                  {interaction.subject}
                                </h4>
                              )}
                              {interaction.notes && (
                                <p className="text-sm text-gray-600 mt-1">{interaction.notes}</p>
                              )}
                            </div>
                          </div>
                        </div>
                      ))
                    )}
                  </div>
                </div>
              </div>
            ) : (
              <div className="bg-white rounded-lg shadow p-12 text-center text-gray-500">
                <Building2 className="w-16 h-16 mx-auto text-gray-300 mb-4" />
                <p>Select a customer to view details</p>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
