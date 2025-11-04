import React, { useState } from 'react';
import { User, Mail, Phone, Building, Edit2, Save } from 'lucide-react';
import { useAuth } from '../context/AuthContext';
import toast from 'react-hot-toast';

const Profile = () => {
  const { user, updateProfile } = useAuth();
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({
    first_name: user?.first_name || '',
    last_name: user?.last_name || '',
    phone: user?.phone || '',
    department: user?.department || '',
  });
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    const success = await updateProfile(formData);
    if (success) {
      setIsEditing(false);
    }

    setLoading(false);
  };

  const handleCancel = () => {
    setFormData({
      first_name: user?.first_name || '',
      last_name: user?.last_name || '',
      phone: user?.phone || '',
      department: user?.department || '',
    });
    setIsEditing(false);
  };

  return (
    <div className="max-w-4xl mx-auto space-y-6">
      {/* Profile Header */}
      <div className="card bg-gradient-to-r from-primary-500 to-primary-600 text-white">
        <div className="flex items-center space-x-6">
          <div className="w-24 h-24 rounded-full bg-white flex items-center justify-center">
            <span className="text-4xl font-bold text-primary-600">
              {user?.first_name?.[0]}{user?.last_name?.[0]}
            </span>
          </div>
          <div className="flex-1">
            <h1 className="text-2xl font-bold mb-1">{user?.full_name}</h1>
            <p className="text-primary-100 mb-2">{user?.email}</p>
            <span className="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-white text-primary-600 capitalize">
              {user?.role}
            </span>
          </div>
          {!isEditing && (
            <button
              onClick={() => setIsEditing(true)}
              className="bg-white text-primary-600 px-4 py-2 rounded-lg hover:bg-primary-50 transition-colors flex items-center space-x-2"
            >
              <Edit2 className="w-4 h-4" />
              <span>Edit Profile</span>
            </button>
          )}
        </div>
      </div>

      {/* Profile Form */}
      <div className="card">
        <h2 className="text-xl font-semibold text-gray-900 mb-6">Profile Information</h2>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                First Name
              </label>
              <div className="relative">
                <User className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                <input
                  type="text"
                  disabled={!isEditing}
                  className="input-field pl-10 disabled:bg-gray-50 disabled:text-gray-500"
                  value={formData.first_name}
                  onChange={(e) => setFormData({ ...formData, first_name: e.target.value })}
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Last Name
              </label>
              <div className="relative">
                <User className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                <input
                  type="text"
                  disabled={!isEditing}
                  className="input-field pl-10 disabled:bg-gray-50 disabled:text-gray-500"
                  value={formData.last_name}
                  onChange={(e) => setFormData({ ...formData, last_name: e.target.value })}
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Email Address
              </label>
              <div className="relative">
                <Mail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                <input
                  type="email"
                  disabled
                  className="input-field pl-10 bg-gray-50 text-gray-500"
                  value={user?.email || ''}
                />
              </div>
              <p className="text-xs text-gray-500 mt-1">Email cannot be changed</p>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Phone Number
              </label>
              <div className="relative">
                <Phone className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                <input
                  type="tel"
                  disabled={!isEditing}
                  className="input-field pl-10 disabled:bg-gray-50 disabled:text-gray-500"
                  placeholder="+1 (555) 123-4567"
                  value={formData.phone}
                  onChange={(e) => setFormData({ ...formData, phone: e.target.value })}
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Department
              </label>
              <div className="relative">
                <Building className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                <input
                  type="text"
                  disabled={!isEditing}
                  className="input-field pl-10 disabled:bg-gray-50 disabled:text-gray-500"
                  placeholder="Engineering"
                  value={formData.department}
                  onChange={(e) => setFormData({ ...formData, department: e.target.value })}
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Role
              </label>
              <input
                type="text"
                disabled
                className="input-field bg-gray-50 text-gray-500 capitalize"
                value={user?.role || ''}
              />
              <p className="text-xs text-gray-500 mt-1">Role is managed by admin</p>
            </div>
          </div>

          {isEditing && (
            <div className="flex items-center justify-end space-x-3 pt-4 border-t">
              <button
                type="button"
                onClick={handleCancel}
                className="btn-secondary"
                disabled={loading}
              >
                Cancel
              </button>
              <button
                type="submit"
                disabled={loading}
                className="btn-primary flex items-center space-x-2"
              >
                {loading ? (
                  <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></div>
                ) : (
                  <>
                    <Save className="w-4 h-4" />
                    <span>Save Changes</span>
                  </>
                )}
              </button>
            </div>
          )}
        </form>
      </div>

      {/* Account Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="card">
          <p className="text-sm text-gray-600 mb-1">Member Since</p>
          <p className="text-lg font-semibold text-gray-900">
            {user?.created_at ? new Date(user.created_at).toLocaleDateString('en-US', {
              month: 'long',
              year: 'numeric'
            }) : 'N/A'}
          </p>
        </div>

        <div className="card">
          <p className="text-sm text-gray-600 mb-1">Account Status</p>
          <p className="text-lg font-semibold text-green-600">
            {user?.is_active ? 'Active' : 'Inactive'}
          </p>
        </div>

        <div className="card">
          <p className="text-sm text-gray-600 mb-1">Email Verified</p>
          <p className="text-lg font-semibold text-gray-900">
            {user?.email_verified ? 'Yes ✓' : 'Not Yet'}
          </p>
        </div>
      </div>
    </div>
  );
};

export default Profile;
