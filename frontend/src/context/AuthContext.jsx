import React, { createContext, useContext, useState, useEffect } from 'react';
import { authAPI, usersAPI } from '../services/api';
import toast from 'react-hot-toast';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const initAuth = async () => {
      const token = localStorage.getItem('access_token');
      if (token) {
        try {
          const response = await usersAPI.getMe();
          setUser(response.data.data);
        } catch (error) {
          localStorage.removeItem('access_token');
          localStorage.removeItem('user');
        }
      }
      setLoading(false);
    };

    initAuth();
  }, []);

  const login = async (email, password) => {
    try {
      const response = await authAPI.login({ email, password });
      const { access_token, user_id, email: userEmail, role } = response.data.data;

      localStorage.setItem('access_token', access_token);

      // Fetch full user profile
      const userResponse = await usersAPI.getMe();
      setUser(userResponse.data.data);

      toast.success('Welcome back!');
      return true;
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Login failed');
      return false;
    }
  };

  const register = async (data) => {
    try {
      await authAPI.register(data);
      toast.success('Registration successful! Please login.');
      return true;
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Registration failed');
      return false;
    }
  };

  const logout = async () => {
    try {
      await authAPI.logout();
    } catch (error) {
      // Ignore error
    } finally {
      localStorage.removeItem('access_token');
      localStorage.removeItem('user');
      setUser(null);
      toast.success('Logged out successfully');
    }
  };

  const updateProfile = async (data) => {
    try {
      const response = await usersAPI.updateMe(data);
      setUser(response.data.data);
      toast.success('Profile updated successfully');
      return true;
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Update failed');
      return false;
    }
  };

  return (
    <AuthContext.Provider value={{ user, loading, login, register, logout, updateProfile }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within AuthProvider');
  }
  return context;
};
