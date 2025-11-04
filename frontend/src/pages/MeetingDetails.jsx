import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Video, Calendar, Clock, Users, ExternalLink, ArrowLeft, AlertCircle } from 'lucide-react';
import { meetingsAPI } from '../services/api';
import { format } from 'date-fns';
import toast from 'react-hot-toast';

const MeetingDetails = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [meeting, setMeeting] = useState(null);
  const [loading, setLoading] = useState(true);
  const [joining, setJoining] = useState(false);

  useEffect(() => {
    loadMeeting();
  }, [id]);

  const loadMeeting = async () => {
    try {
      const response = await meetingsAPI.get(id);
      setMeeting(response.data.data);
    } catch (error) {
      toast.error('Failed to load meeting details');
      navigate('/meetings');
    } finally {
      setLoading(false);
    }
  };

  const handleJoinMeeting = async () => {
    setJoining(true);
    try {
      const response = await meetingsAPI.join(id);
      const { room_url, user_name, user_email } = response.data.data;

      // Open Jitsi meeting in new window without JWT
      // User info will be displayed when they join
      window.open(room_url, '_blank', 'width=1200,height=800');

      toast.success('Joined meeting successfully!');
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Failed to join meeting');
    } finally {
      setJoining(false);
    }
  };

  const getStatusColor = (status) => {
    const colors = {
      scheduled: 'bg-blue-100 text-blue-700 border-blue-200',
      ongoing: 'bg-green-100 text-green-700 border-green-200',
      completed: 'bg-gray-100 text-gray-700 border-gray-200',
      cancelled: 'bg-red-100 text-red-700 border-red-200',
    };
    return colors[status] || 'bg-gray-100 text-gray-700 border-gray-200';
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600"></div>
      </div>
    );
  }

  if (!meeting) {
    return (
      <div className="card text-center py-12">
        <AlertCircle className="w-12 h-12 text-gray-400 mx-auto mb-4" />
        <p className="text-gray-600">Meeting not found</p>
      </div>
    );
  }

  const canJoin = meeting.status === 'scheduled' || meeting.status === 'ongoing';

  return (
    <div className="max-w-4xl mx-auto space-y-6">
      {/* Back Button */}
      <button
        onClick={() => navigate('/meetings')}
        className="flex items-center space-x-2 text-gray-600 hover:text-gray-900 transition-colors"
      >
        <ArrowLeft className="w-5 h-5" />
        <span>Back to Meetings</span>
      </button>

      {/* Meeting Header */}
      <div className="card">
        <div className="flex items-start justify-between mb-6">
          <div className="flex items-start space-x-4">
            <div className="p-3 bg-primary-100 rounded-lg">
              <Video className="w-8 h-8 text-primary-600" />
            </div>
            <div>
              <h1 className="text-2xl font-bold text-gray-900 mb-2">{meeting.title}</h1>
              <span className={`inline-block px-3 py-1 rounded-full text-sm font-medium border capitalize ${getStatusColor(meeting.status)}`}>
                {meeting.status}
              </span>
            </div>
          </div>

          {canJoin && (
            <button
              onClick={handleJoinMeeting}
              disabled={joining}
              className="btn-primary flex items-center space-x-2"
            >
              {joining ? (
                <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></div>
              ) : (
                <>
                  <ExternalLink className="w-5 h-5" />
                  <span>Join Meeting</span>
                </>
              )}
            </button>
          )}
        </div>

        {/* Meeting Info Grid */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6 pb-6 border-b">
          <div className="flex items-center space-x-3">
            <div className="p-2 bg-gray-100 rounded-lg">
              <Calendar className="w-5 h-5 text-gray-600" />
            </div>
            <div>
              <p className="text-xs text-gray-500">Date</p>
              <p className="font-medium text-gray-900">
                {format(new Date(meeting.start_time), 'MMMM dd, yyyy')}
              </p>
            </div>
          </div>

          <div className="flex items-center space-x-3">
            <div className="p-2 bg-gray-100 rounded-lg">
              <Clock className="w-5 h-5 text-gray-600" />
            </div>
            <div>
              <p className="text-xs text-gray-500">Time</p>
              <p className="font-medium text-gray-900">
                {format(new Date(meeting.start_time), 'HH:mm')}
              </p>
            </div>
          </div>

          <div className="flex items-center space-x-3">
            <div className="p-2 bg-gray-100 rounded-lg">
              <Users className="w-5 h-5 text-gray-600" />
            </div>
            <div>
              <p className="text-xs text-gray-500">Max Participants</p>
              <p className="font-medium text-gray-900">{meeting.max_participants}</p>
            </div>
          </div>
        </div>

        {/* Description */}
        {meeting.description && (
          <div>
            <h3 className="font-semibold text-gray-900 mb-2">Description</h3>
            <p className="text-gray-600 whitespace-pre-wrap">{meeting.description}</p>
          </div>
        )}
      </div>

      {/* Meeting Details */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="card">
          <h3 className="font-semibold text-gray-900 mb-4">Meeting Information</h3>
          <div className="space-y-3 text-sm">
            <div className="flex justify-between">
              <span className="text-gray-600">Room ID:</span>
              <span className="font-mono text-gray-900">{meeting.room_id}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600">Organizer ID:</span>
              <span className="font-mono text-gray-900">{meeting.organizer_id.slice(0, 8)}...</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600">Created:</span>
              <span className="text-gray-900">
                {format(new Date(meeting.created_at), 'MMM dd, HH:mm')}
              </span>
            </div>
            {meeting.jitsi_room_url && (
              <div className="flex justify-between">
                <span className="text-gray-600">Room URL:</span>
                <span className="text-primary-600 truncate max-w-xs">Active</span>
              </div>
            )}
          </div>
        </div>

        {meeting.status === 'completed' && meeting.recording_url && (
          <div className="card">
            <h3 className="font-semibold text-gray-900 mb-4">Recording</h3>
            <a
              href={meeting.recording_url}
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center justify-center space-x-2 w-full py-3 bg-primary-50 text-primary-700 rounded-lg hover:bg-primary-100 transition-colors"
            >
              <ExternalLink className="w-4 h-4" />
              <span>View Recording</span>
            </a>
          </div>
        )}

        {canJoin && (
          <div className="card bg-green-50 border-green-200">
            <h3 className="font-semibold text-green-900 mb-2">Ready to join?</h3>
            <p className="text-sm text-green-700 mb-4">
              Click the "Join Meeting" button to start your video conference. Make sure your camera and microphone are working.
            </p>
            <button
              onClick={handleJoinMeeting}
              disabled={joining}
              className="w-full btn-primary flex items-center justify-center space-x-2"
            >
              {joining ? (
                <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></div>
              ) : (
                <>
                  <ExternalLink className="w-5 h-5" />
                  <span>Join Now</span>
                </>
              )}
            </button>
          </div>
        )}
      </div>

      {/* Info Alert */}
      <div className="card bg-blue-50 border-blue-200">
        <div className="flex items-start space-x-3">
          <AlertCircle className="w-5 h-5 text-blue-600 mt-0.5" />
          <div>
            <h3 className="font-medium text-blue-900 mb-1">Meeting Tips</h3>
            <ul className="text-sm text-blue-700 space-y-1 list-disc list-inside">
              <li>Join a few minutes early to test your audio and video</li>
              <li>Use a stable internet connection for best quality</li>
              <li>Mute your microphone when not speaking</li>
              <li>Use screen sharing to present documents</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MeetingDetails;
