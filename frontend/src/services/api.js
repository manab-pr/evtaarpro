import axios from 'axios';

const API_BASE_URL = '/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('access_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor to handle errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('access_token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Auth APIs
export const authAPI = {
  register: (data) => api.post('/auth/register', data),
  login: (data) => api.post('/auth/login', data),
  logout: () => api.post('/auth/logout'),
};

// Users APIs
export const usersAPI = {
  getMe: () => api.get('/users/me'),
  updateMe: (data) => api.put('/users/me', data),
  getUsers: (params) => api.get('/users', { params }),
  getUser: (id) => api.get(`/users/${id}`),
};

// Meetings APIs
export const meetingsAPI = {
  create: (data) => api.post('/meetings', data),
  list: (params) => api.get('/meetings', { params }),
  get: (id) => api.get(`/meetings/${id}`),
  join: (id) => api.post(`/meetings/${id}/join`),
};

export default api;
