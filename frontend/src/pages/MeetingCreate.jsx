import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Video, Calendar, FileText, Users, ArrowLeft } from 'lucide-react';
import { meetingsAPI } from '../services/api';
import toast from 'react-hot-toast';

const MeetingCreate = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    start_time: '',
    max_participants: 50,
  });
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
      const response = await meetingsAPI.create({
        ...formData,
        start_time: new Date(formData.start_time).toISOString(),
      });

      toast.success('Meeting created successfully!');
      navigate(`/meetings/${response.data.data.id}`);
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Failed to create meeting');
    } finally {
      setLoading(false);
    }
  };

  // Get current datetime for min attribute
  const now = new Date();
  now.setMinutes(now.getMinutes() - now.getTimezoneOffset());
  const minDateTime = now.toISOString().slice(0, 16);

  return (
    <div className="max-w-3xl mx-auto space-y-6">
      {/* Back Button */}
      <button
        onClick={() => navigate('/meetings')}
        className="flex items-center space-x-2 text-gray-600 hover:text-gray-900 transition-colors"
      >
        <ArrowLeft className="w-5 h-5" />
        <span>Back to Meetings</span>
      </button>

      {/* Header */}
      <div className="card bg-gradient-to-r from-primary-500 to-primary-600 text-white">
        <div className="flex items-center space-x-4">
          <div className="p-3 bg-white rounded-lg">
            <Video className="w-8 h-8 text-primary-600" />
          </div>
          <div>
            <h1 className="text-2xl font-bold">Create New Meeting</h1>
            <p className="text-primary-100 mt-1">Schedule a video meeting with your team</p>
          </div>
        </div>
      </div>

      {/* Form */}
      <div className="card">
        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Meeting Title */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Meeting Title *
            </label>
            <div className="relative">
              <FileText className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
              <input
                type="text"
                required
                className="input-field pl-10"
                placeholder="e.g., Team Standup, Client Review, etc."
                value={formData.title}
                onChange={(e) => setFormData({ ...formData, title: e.target.value })}
              />
            </div>
          </div>

          {/* Description */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Description
            </label>
            <textarea
              rows={4}
              className="input-field resize-none"
              placeholder="Add meeting agenda, topics to discuss, or any important notes..."
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
            />
          </div>

          {/* Date and Time */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Start Date & Time *
            </label>
            <div className="relative">
              <Calendar className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
              <input
                type="datetime-local"
                required
                min={minDateTime}
                className="input-field pl-10"
                value={formData.start_time}
                onChange={(e) => setFormData({ ...formData, start_time: e.target.value })}
              />
            </div>
            <p className="text-xs text-gray-500 mt-1">
              Select when you want to start the meeting
            </p>
          </div>

          {/* Max Participants */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Maximum Participants
            </label>
            <div className="relative">
              <Users className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
              <input
                type="number"
                min={2}
                max={500}
                className="input-field pl-10"
                value={formData.max_participants}
                onChange={(e) => setFormData({ ...formData, max_participants: parseInt(e.target.value) })}
              />
            </div>
            <p className="text-xs text-gray-500 mt-1">
              Maximum number of people who can join (2-500)
            </p>
          </div>

          {/* Actions */}
          <div className="flex items-center justify-end space-x-3 pt-6 border-t">
            <button
              type="button"
              onClick={() => navigate('/meetings')}
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
                  <Video className="w-5 h-5" />
                  <span>Create Meeting</span>
                </>
              )}
            </button>
          </div>
        </form>
      </div>

      {/* Info Card */}
      <div className="card bg-blue-50 border-blue-200">
        <div className="flex items-start space-x-3">
          <div className="p-2 bg-blue-100 rounded-lg">
            <Video className="w-5 h-5 text-blue-600" />
          </div>
          <div>
            <h3 className="font-medium text-blue-900 mb-1">About Video Meetings</h3>
            <p className="text-sm text-blue-700">
              Once created, you and your team members can join the meeting from the meetings page.
              The meeting uses Jitsi for secure, high-quality video conferencing with features like
              screen sharing, recording, and more.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MeetingCreate;
