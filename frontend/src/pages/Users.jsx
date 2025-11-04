import React, { useEffect, useState } from 'react';
import { Search, Users as UsersIcon, Mail, Phone } from 'lucide-react';
import { usersAPI } from '../services/api';
import toast from 'react-hot-toast';

const Users = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [page, setPage] = useState(1);
  const [pagination, setPagination] = useState(null);

  useEffect(() => {
    loadUsers();
  }, [page]);

  const loadUsers = async () => {
    try {
      setLoading(true);
      const response = await usersAPI.getUsers({ page, page_size: 12, search: searchQuery });
      setUsers(response.data.data || []);
      setPagination(response.data.pagination);
    } catch (error) {
      toast.error('Failed to load users');
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = (e) => {
    e.preventDefault();
    setPage(1);
    loadUsers();
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Team Members</h1>
          <p className="text-gray-600 mt-1">Manage and view all team members</p>
        </div>
        <div className="flex items-center space-x-2 text-gray-600">
          <UsersIcon className="w-5 h-5" />
          <span className="font-medium">{pagination?.total_items || 0} Members</span>
        </div>
      </div>

      {/* Search */}
      <div className="card">
        <form onSubmit={handleSearch} className="flex items-center space-x-4">
          <div className="flex-1 relative">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
            <input
              type="text"
              placeholder="Search by name or email..."
              className="input-field pl-10"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
          </div>
          <button type="submit" className="btn-primary">
            Search
          </button>
        </form>
      </div>

      {/* Users Grid */}
      {loading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[...Array(6)].map((_, i) => (
            <div key={i} className="card animate-pulse">
              <div className="flex items-center space-x-4">
                <div className="w-12 h-12 bg-gray-200 rounded-full"></div>
                <div className="flex-1 space-y-2">
                  <div className="h-4 bg-gray-200 rounded w-3/4"></div>
                  <div className="h-3 bg-gray-200 rounded w-1/2"></div>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : users.length === 0 ? (
        <div className="card text-center py-12">
          <UsersIcon className="w-12 h-12 text-gray-400 mx-auto mb-4" />
          <p className="text-gray-600">No users found</p>
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {users.map((user) => (
              <div key={user.id} className="card hover:shadow-md transition-shadow">
                <div className="flex items-start space-x-4">
                  <div className="w-12 h-12 rounded-full bg-primary-100 flex items-center justify-center flex-shrink-0">
                    <span className="text-primary-600 font-semibold">
                      {user.first_name?.[0]}{user.last_name?.[0]}
                    </span>
                  </div>

                  <div className="flex-1 min-w-0">
                    <h3 className="font-semibold text-gray-900 truncate">
                      {user.full_name}
                    </h3>
                    <p className="text-sm text-gray-600 truncate flex items-center mt-1">
                      <Mail className="w-3 h-3 mr-1" />
                      {user.email}
                    </p>
                    {user.phone && (
                      <p className="text-sm text-gray-600 truncate flex items-center mt-1">
                        <Phone className="w-3 h-3 mr-1" />
                        {user.phone}
                      </p>
                    )}
                    <div className="flex items-center space-x-2 mt-3">
                      <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-primary-100 text-primary-800 capitalize">
                        {user.role}
                      </span>
                      {user.department && (
                        <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-800">
                          {user.department}
                        </span>
                      )}
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>

          {/* Pagination */}
          {pagination && pagination.total_pages > 1 && (
            <div className="flex items-center justify-center space-x-2">
              <button
                onClick={() => setPage(p => Math.max(1, p - 1))}
                disabled={page === 1}
                className="btn-secondary disabled:opacity-50"
              >
                Previous
              </button>
              <span className="text-sm text-gray-600">
                Page {page} of {pagination.total_pages}
              </span>
              <button
                onClick={() => setPage(p => Math.min(pagination.total_pages, p + 1))}
                disabled={page === pagination.total_pages}
                className="btn-secondary disabled:opacity-50"
              >
                Next
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
};

export default Users;
