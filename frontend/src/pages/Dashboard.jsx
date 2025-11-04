import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Video, Users, Calendar, TrendingUp, Plus } from 'lucide-react';
import { useAuth } from '../context/AuthContext';
import { meetingsAPI, usersAPI } from '../services/api';
import { format } from 'date-fns';

const Dashboard = () => {
  const { user } = useAuth();
  const [stats, setStats] = useState({
    totalMeetings: 0,
    upcomingMeetings: 0,
    totalUsers: 0,
  });
  const [recentMeetings, setRecentMeetings] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadDashboardData();
  }, []);

  const loadDashboardData = async () => {
    try {
      const [meetingsRes, usersRes] = await Promise.all([
        meetingsAPI.list({ page: 1, page_size: 5 }),
        usersAPI.getUsers({ page: 1, page_size: 1 }),
      ]);

      setRecentMeetings(meetingsRes.data.data || []);
      setStats({
        totalMeetings: meetingsRes.data.pagination?.total_items || 0,
        upcomingMeetings: meetingsRes.data.data?.filter(m => m.status === 'scheduled').length || 0,
        totalUsers: usersRes.data.pagination?.total_items || 0,
      });
    } catch (error) {
      console.error('Failed to load dashboard data:', error);
    } finally {
      setLoading(false);
    }
  };

  const statCards = [
    {
      title: 'Total Meetings',
      value: stats.totalMeetings,
      icon: Video,
      color: 'bg-blue-500',
      bgColor: 'bg-blue-50',
      textColor: 'text-blue-600',
    },
    {
      title: 'Upcoming Meetings',
      value: stats.upcomingMeetings,
      icon: Calendar,
      color: 'bg-green-500',
      bgColor: 'bg-green-50',
      textColor: 'text-green-600',
    },
    {
      title: 'Team Members',
      value: stats.totalUsers,
      icon: Users,
      color: 'bg-purple-500',
      bgColor: 'bg-purple-50',
      textColor: 'text-purple-600',
    },
    {
      title: 'Active Projects',
      value: '12',
      icon: TrendingUp,
      color: 'bg-orange-500',
      bgColor: 'bg-orange-50',
      textColor: 'text-orange-600',
    },
  ];

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600"></div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Welcome Section */}
      <div className="card bg-gradient-to-r from-primary-500 to-primary-600 text-white">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold mb-2">
              Welcome back, {user?.first_name}! 👋
            </h1>
            <p className="text-primary-100">
              Here's what's happening with your projects today.
            </p>
          </div>
          <Link
            to="/meetings/new"
            className="hidden md:flex items-center space-x-2 bg-white text-primary-600 px-4 py-2 rounded-lg hover:bg-primary-50 transition-colors"
          >
            <Plus className="w-5 h-5" />
            <span>New Meeting</span>
          </Link>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {statCards.map((stat, index) => {
          const Icon = stat.icon;
          return (
            <div
              key={index}
              className="card hover:shadow-md transition-shadow cursor-pointer"
            >
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-gray-600 mb-1">
                    {stat.title}
                  </p>
                  <p className="text-2xl font-bold text-gray-900">
                    {stat.value}
                  </p>
                </div>
                <div className={`${stat.bgColor} p-3 rounded-lg`}>
                  <Icon className={`w-6 h-6 ${stat.textColor}`} />
                </div>
              </div>
            </div>
          );
        })}
      </div>

      {/* Recent Meetings */}
      <div className="card">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-lg font-semibold text-gray-900">Recent Meetings</h2>
          <Link
            to="/meetings"
            className="text-sm text-primary-600 hover:text-primary-700 font-medium"
          >
            View all →
          </Link>
        </div>

        {recentMeetings.length === 0 ? (
          <div className="text-center py-12">
            <Video className="w-12 h-12 text-gray-400 mx-auto mb-4" />
            <p className="text-gray-600 mb-4">No meetings yet</p>
            <Link to="/meetings/new" className="btn-primary inline-block">
              Create Your First Meeting
            </Link>
          </div>
        ) : (
          <div className="space-y-3">
            {recentMeetings.map((meeting) => (
              <Link
                key={meeting.id}
                to={`/meetings/${meeting.id}`}
                className="block p-4 rounded-lg border border-gray-200 hover:border-primary-300 hover:bg-primary-50 transition-all"
              >
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-3">
                    <div className="p-2 bg-primary-100 rounded-lg">
                      <Video className="w-5 h-5 text-primary-600" />
                    </div>
                    <div>
                      <h3 className="font-medium text-gray-900">{meeting.title}</h3>
                      <p className="text-sm text-gray-500">
                        {format(new Date(meeting.start_time), 'MMM dd, yyyy • HH:mm')}
                      </p>
                    </div>
                  </div>
                  <span className={`
                    px-3 py-1 rounded-full text-xs font-medium capitalize
                    ${meeting.status === 'scheduled' ? 'bg-blue-100 text-blue-700' : ''}
                    ${meeting.status === 'ongoing' ? 'bg-green-100 text-green-700' : ''}
                    ${meeting.status === 'completed' ? 'bg-gray-100 text-gray-700' : ''}
                  `}>
                    {meeting.status}
                  </span>
                </div>
              </Link>
            ))}
          </div>
        )}
      </div>

      {/* Quick Actions */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <Link
          to="/meetings/new"
          className="card hover:shadow-md transition-shadow group"
        >
          <div className="flex items-center space-x-4">
            <div className="p-3 bg-primary-100 rounded-lg group-hover:bg-primary-200 transition-colors">
              <Plus className="w-6 h-6 text-primary-600" />
            </div>
            <div>
              <h3 className="font-medium text-gray-900">Schedule Meeting</h3>
              <p className="text-sm text-gray-500">Create a new video meeting</p>
            </div>
          </div>
        </Link>

        <Link
          to="/users"
          className="card hover:shadow-md transition-shadow group"
        >
          <div className="flex items-center space-x-4">
            <div className="p-3 bg-purple-100 rounded-lg group-hover:bg-purple-200 transition-colors">
              <Users className="w-6 h-6 text-purple-600" />
            </div>
            <div>
              <h3 className="font-medium text-gray-900">View Team</h3>
              <p className="text-sm text-gray-500">See all team members</p>
            </div>
          </div>
        </Link>

        <Link
          to="/profile"
          className="card hover:shadow-md transition-shadow group"
        >
          <div className="flex items-center space-x-4">
            <div className="p-3 bg-green-100 rounded-lg group-hover:bg-green-200 transition-colors">
              <Calendar className="w-6 h-6 text-green-600" />
            </div>
            <div>
              <h3 className="font-medium text-gray-900">My Profile</h3>
              <p className="text-sm text-gray-500">Update your information</p>
            </div>
          </div>
        </Link>
      </div>
    </div>
  );
};

export default Dashboard;
