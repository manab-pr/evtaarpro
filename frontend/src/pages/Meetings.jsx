import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Video, Plus, Calendar, Clock, Users } from 'lucide-react';
import { meetingsAPI } from '../services/api';
import { format } from 'date-fns';
import toast from 'react-hot-toast';

const Meetings = () => {
  const [meetings, setMeetings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [pagination, setPagination] = useState(null);

  useEffect(() => {
    loadMeetings();
  }, [page]);

  const loadMeetings = async () => {
    try {
      setLoading(true);
      const response = await meetingsAPI.list({ page, page_size: 10 });
      setMeetings(response.data.data || []);
      setPagination(response.data.pagination);
    } catch (error) {
      toast.error('Failed to load meetings');
    } finally {
      setLoading(false);
    }
  };

  const getStatusColor = (status) => {
    const colors = {
      scheduled: 'bg-blue-100 text-blue-700',
      ongoing: 'bg-green-100 text-green-700',
      completed: 'bg-gray-100 text-gray-700',
      cancelled: 'bg-red-100 text-red-700',
    };
    return colors[status] || 'bg-gray-100 text-gray-700';
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Meetings</h1>
          <p className="text-gray-600 mt-1">Manage and join your meetings</p>
        </div>
        <Link to="/meetings/new" className="btn-primary flex items-center space-x-2">
          <Plus className="w-5 h-5" />
          <span>New Meeting</span>
        </Link>
      </div>

      {/* Meetings List */}
      {loading ? (
        <div className="space-y-4">
          {[...Array(5)].map((_, i) => (
            <div key={i} className="card animate-pulse">
              <div className="flex items-center space-x-4">
                <div className="w-12 h-12 bg-gray-200 rounded-lg"></div>
                <div className="flex-1 space-y-2">
                  <div className="h-4 bg-gray-200 rounded w-1/3"></div>
                  <div className="h-3 bg-gray-200 rounded w-1/4"></div>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : meetings.length === 0 ? (
        <div className="card text-center py-16">
          <Video className="w-16 h-16 text-gray-400 mx-auto mb-4" />
          <h3 className="text-lg font-medium text-gray-900 mb-2">No meetings yet</h3>
          <p className="text-gray-600 mb-6">Get started by creating your first meeting</p>
          <Link to="/meetings/new" className="btn-primary inline-flex items-center space-x-2">
            <Plus className="w-5 h-5" />
            <span>Create Meeting</span>
          </Link>
        </div>
      ) : (
        <>
          <div className="space-y-4">
            {meetings.map((meeting) => (
              <Link
                key={meeting.id}
                to={`/meetings/${meeting.id}`}
                className="card hover:shadow-md transition-all group"
              >
                <div className="flex items-start justify-between">
                  <div className="flex items-start space-x-4 flex-1">
                    <div className="p-3 bg-primary-100 rounded-lg group-hover:bg-primary-200 transition-colors">
                      <Video className="w-6 h-6 text-primary-600" />
                    </div>

                    <div className="flex-1">
                      <h3 className="text-lg font-semibold text-gray-900 mb-1">
                        {meeting.title}
                      </h3>
                      {meeting.description && (
                        <p className="text-sm text-gray-600 mb-3 line-clamp-2">
                          {meeting.description}
                        </p>
                      )}

                      <div className="flex flex-wrap items-center gap-4 text-sm text-gray-600">
                        <div className="flex items-center">
                          <Calendar className="w-4 h-4 mr-1" />
                          {format(new Date(meeting.start_time), 'MMM dd, yyyy')}
                        </div>
                        <div className="flex items-center">
                          <Clock className="w-4 h-4 mr-1" />
                          {format(new Date(meeting.start_time), 'HH:mm')}
                        </div>
                        <div className="flex items-center">
                          <Users className="w-4 h-4 mr-1" />
                          Max {meeting.max_participants} participants
                        </div>
                      </div>
                    </div>
                  </div>

                  <div className="flex flex-col items-end space-y-2">
                    <span className={`px-3 py-1 rounded-full text-xs font-medium capitalize ${getStatusColor(meeting.status)}`}>
                      {meeting.status}
                    </span>
                    {(meeting.status === 'scheduled' || meeting.status === 'ongoing') && (
                      <span className="text-xs text-primary-600 font-medium group-hover:text-primary-700">
                        Click to join →
                      </span>
                    )}
                  </div>
                </div>
              </Link>
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

export default Meetings;
