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

// CRM APIs
export const crmAPI = {
  // Customers
  createCustomer: (data) => api.post('/crm/customers', data),
  listCustomers: (params) => api.get('/crm/customers', { params }),
  getCustomer: (id) => api.get(`/crm/customers/${id}`),
  updateCustomer: (id, data) => api.put(`/crm/customers/${id}`, data),
  deleteCustomer: (id) => api.delete(`/crm/customers/${id}`),

  // Interactions
  addInteraction: (customerId, data) => api.post(`/crm/customers/${customerId}/interactions`, data),
  getInteractions: (customerId) => api.get(`/crm/customers/${customerId}/interactions`),
};

// Payroll APIs
export const payrollAPI = {
  // Employees
  createEmployee: (data) => api.post('/payroll/employees', data),
  listEmployees: (params) => api.get('/payroll/employees', { params }),
  getEmployee: (id) => api.get(`/payroll/employees/${id}`),

  // Attendance
  markAttendance: (data) => api.post('/payroll/attendance', data),
  getAttendance: (params) => api.get('/payroll/attendance', { params }),

  // Payroll Records
  generatePayroll: (data) => api.post('/payroll/records', data),
  listPayrollRecords: (params) => api.get('/payroll/records', { params }),
  getPayrollRecord: (id) => api.get(`/payroll/records/${id}`),
};

// Notifications APIs
export const notificationsAPI = {
  list: (params) => api.get('/notifications', { params }),
  getUnreadCount: () => api.get('/notifications/unread-count'),
  markAsRead: (id) => api.put(`/notifications/${id}/read`),
  markAllAsRead: () => api.post('/notifications/read-all'),
  delete: (id) => api.delete(`/notifications/${id}`),
};

export default api;
