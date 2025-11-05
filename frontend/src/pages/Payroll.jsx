import React, { useState, useEffect } from 'react';
import { Users, Plus, Calendar, DollarSign, Clock, TrendingUp } from 'lucide-react';
import { payrollAPI } from '../services/api';
import toast from 'react-hot-toast';

const Payroll = () => {
  const [activeTab, setActiveTab] = useState('employees');
  const [employees, setEmployees] = useState([]);
  const [attendance, setAttendance] = useState([]);
  const [payrollRecords, setPayrollRecords] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showAddEmployeeModal, setShowAddEmployeeModal] = useState(false);
  const [showAttendanceModal, setShowAttendanceModal] = useState(false);
  const [showPayrollModal, setShowPayrollModal] = useState(false);

  const [employeeForm, setEmployeeForm] = useState({
    employee_code: '',
    department: '',
    designation: '',
    joining_date: '',
    salary_amount: '',
  });

  const [attendanceForm, setAttendanceForm] = useState({
    employee_id: '',
    date: '',
    check_in: '',
    check_out: '',
    status: 'present',
    notes: '',
  });

  const [payrollForm, setPayrollForm] = useState({
    employee_id: '',
    month: new Date().getMonth() + 1,
    year: new Date().getFullYear(),
    allowances: '',
    deductions: '',
    notes: '',
  });

  useEffect(() => {
    loadData();
  }, [activeTab]);

  const loadData = async () => {
    setLoading(true);
    try {
      if (activeTab === 'employees') {
        const response = await payrollAPI.listEmployees({ page: 1, page_size: 50 });
        setEmployees(response.data.data || []);
      } else if (activeTab === 'attendance') {
        const response = await payrollAPI.getAttendance({ page: 1, page_size: 50 });
        setAttendance(response.data.data || []);
      } else if (activeTab === 'payroll') {
        const response = await payrollAPI.listPayrollRecords({ page: 1, page_size: 50 });
        setPayrollRecords(response.data.data || []);
      }
    } catch (error) {
      toast.error('Failed to load data');
    } finally {
      setLoading(false);
    }
  };

  const handleAddEmployee = async (e) => {
    e.preventDefault();
    try {
      await payrollAPI.createEmployee({
        ...employeeForm,
        salary_amount: parseFloat(employeeForm.salary_amount),
      });
      toast.success('Employee added successfully!');
      setShowAddEmployeeModal(false);
      setEmployeeForm({
        employee_code: '',
        department: '',
        designation: '',
        joining_date: '',
        salary_amount: '',
      });
      loadData();
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Failed to add employee');
    }
  };

  const handleMarkAttendance = async (e) => {
    e.preventDefault();
    try {
      await payrollAPI.markAttendance({
        ...attendanceForm,
        check_in: attendanceForm.check_in ? `${attendanceForm.date}T${attendanceForm.check_in}:00Z` : null,
        check_out: attendanceForm.check_out ? `${attendanceForm.date}T${attendanceForm.check_out}:00Z` : null,
      });
      toast.success('Attendance marked successfully!');
      setShowAttendanceModal(false);
      setAttendanceForm({
        employee_id: '',
        date: '',
        check_in: '',
        check_out: '',
        status: 'present',
        notes: '',
      });
      loadData();
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Failed to mark attendance');
    }
  };

  const handleGeneratePayroll = async (e) => {
    e.preventDefault();
    try {
      await payrollAPI.generatePayroll({
        ...payrollForm,
        allowances: parseFloat(payrollForm.allowances || 0),
        deductions: parseFloat(payrollForm.deductions || 0),
      });
      toast.success('Payroll generated successfully!');
      setShowPayrollModal(false);
      setPayrollForm({
        employee_id: '',
        month: new Date().getMonth() + 1,
        year: new Date().getFullYear(),
        allowances: '',
        deductions: '',
        notes: '',
      });
      loadData();
    } catch (error) {
      toast.error(error.response?.data?.error?.message || 'Failed to generate payroll');
    }
  };

  const formatCurrency = (amount) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
    }).format(amount);
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600"></div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Payroll Management</h1>
          <p className="text-gray-600 mt-1">Manage employees, attendance, and payroll</p>
        </div>
        <button
          onClick={() => {
            if (activeTab === 'employees') setShowAddEmployeeModal(true);
            else if (activeTab === 'attendance') setShowAttendanceModal(true);
            else if (activeTab === 'payroll') setShowPayrollModal(true);
          }}
          className="btn-primary flex items-center space-x-2"
        >
          <Plus className="w-5 h-5" />
          <span>
            {activeTab === 'employees' && 'Add Employee'}
            {activeTab === 'attendance' && 'Mark Attendance'}
            {activeTab === 'payroll' && 'Generate Payroll'}
          </span>
        </button>
      </div>

      {/* Tabs */}
      <div className="card">
        <div className="flex space-x-4 border-b">
          <button
            onClick={() => setActiveTab('employees')}
            className={`px-4 py-2 font-medium transition-colors ${
              activeTab === 'employees'
                ? 'border-b-2 border-primary-600 text-primary-600'
                : 'text-gray-600 hover:text-gray-900'
            }`}
          >
            Employees
          </button>
          <button
            onClick={() => setActiveTab('attendance')}
            className={`px-4 py-2 font-medium transition-colors ${
              activeTab === 'attendance'
                ? 'border-b-2 border-primary-600 text-primary-600'
                : 'text-gray-600 hover:text-gray-900'
            }`}
          >
            Attendance
          </button>
          <button
            onClick={() => setActiveTab('payroll')}
            className={`px-4 py-2 font-medium transition-colors ${
              activeTab === 'payroll'
                ? 'border-b-2 border-primary-600 text-primary-600'
                : 'text-gray-600 hover:text-gray-900'
            }`}
          >
            Payroll Records
          </button>
        </div>
      </div>

      {/* Stats */}
      {activeTab === 'employees' && (
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
          <div className="card">
            <div className="text-sm text-gray-600">Total Employees</div>
            <div className="text-2xl font-bold text-gray-900 mt-1">{employees.length}</div>
          </div>
          <div className="card">
            <div className="text-sm text-gray-600">Active Employees</div>
            <div className="text-2xl font-bold text-green-600 mt-1">
              {employees.filter(e => e.is_active).length}
            </div>
          </div>
          <div className="card">
            <div className="text-sm text-gray-600">Total Salary</div>
            <div className="text-2xl font-bold text-gray-900 mt-1">
              {formatCurrency(employees.reduce((sum, e) => sum + e.salary_amount, 0))}
            </div>
          </div>
          <div className="card">
            <div className="text-sm text-gray-600">Avg. Salary</div>
            <div className="text-2xl font-bold text-gray-900 mt-1">
              {formatCurrency(employees.length > 0 ? employees.reduce((sum, e) => sum + e.salary_amount, 0) / employees.length : 0)}
            </div>
          </div>
        </div>
      )}

      {/* Content */}
      <div className="card">
        <div className="overflow-x-auto">
          {activeTab === 'employees' && (
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Code</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Department</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Designation</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Joining Date</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Salary</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {employees.length === 0 ? (
                  <tr>
                    <td colSpan="6" className="px-6 py-8 text-center text-gray-500">
                      No employees found. Click "Add Employee" to get started.
                    </td>
                  </tr>
                ) : (
                  employees.map((employee) => (
                    <tr key={employee.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4">
                        <div className="flex items-center space-x-3">
                          <div className="w-10 h-10 rounded-full bg-primary-100 flex items-center justify-center">
                            <Users className="w-5 h-5 text-primary-600" />
                          </div>
                          <div className="font-medium text-gray-900">{employee.employee_code}</div>
                        </div>
                      </td>
                      <td className="px-6 py-4 text-gray-900">{employee.department}</td>
                      <td className="px-6 py-4 text-gray-900">{employee.designation}</td>
                      <td className="px-6 py-4 text-gray-600">{formatDate(employee.joining_date)}</td>
                      <td className="px-6 py-4 text-gray-900 font-medium">{formatCurrency(employee.salary_amount)}</td>
                      <td className="px-6 py-4">
                        <span
                          className={`inline-block px-3 py-1 rounded-full text-xs font-medium ${
                            employee.is_active
                              ? 'bg-green-100 text-green-700 border border-green-200'
                              : 'bg-gray-100 text-gray-700 border border-gray-200'
                          }`}
                        >
                          {employee.is_active ? 'Active' : 'Inactive'}
                        </span>
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          )}

          {activeTab === 'attendance' && (
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Employee</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Date</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Check In</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Check Out</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Hours</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {attendance.length === 0 ? (
                  <tr>
                    <td colSpan="6" className="px-6 py-8 text-center text-gray-500">
                      No attendance records found. Click "Mark Attendance" to get started.
                    </td>
                  </tr>
                ) : (
                  attendance.map((record) => (
                    <tr key={record.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4">
                        <div className="flex items-center space-x-3">
                          <Clock className="w-5 h-5 text-gray-400" />
                          <div className="font-medium text-gray-900">{record.employee_id}</div>
                        </div>
                      </td>
                      <td className="px-6 py-4 text-gray-900">{formatDate(record.date)}</td>
                      <td className="px-6 py-4 text-gray-600">
                        {record.check_in ? new Date(record.check_in).toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' }) : '-'}
                      </td>
                      <td className="px-6 py-4 text-gray-600">
                        {record.check_out ? new Date(record.check_out).toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' }) : '-'}
                      </td>
                      <td className="px-6 py-4 text-gray-900">{record.hours_worked?.toFixed(2) || '0.00'} hrs</td>
                      <td className="px-6 py-4">
                        <span
                          className={`inline-block px-3 py-1 rounded-full text-xs font-medium capitalize ${
                            record.status === 'present'
                              ? 'bg-green-100 text-green-700 border border-green-200'
                              : record.status === 'absent'
                              ? 'bg-red-100 text-red-700 border border-red-200'
                              : 'bg-yellow-100 text-yellow-700 border border-yellow-200'
                          }`}
                        >
                          {record.status}
                        </span>
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          )}

          {activeTab === 'payroll' && (
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Employee</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Period</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Basic Salary</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Allowances</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Deductions</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Net Salary</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {payrollRecords.length === 0 ? (
                  <tr>
                    <td colSpan="7" className="px-6 py-8 text-center text-gray-500">
                      No payroll records found. Click "Generate Payroll" to get started.
                    </td>
                  </tr>
                ) : (
                  payrollRecords.map((record) => (
                    <tr key={record.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4">
                        <div className="flex items-center space-x-3">
                          <DollarSign className="w-5 h-5 text-gray-400" />
                          <div className="font-medium text-gray-900">{record.employee_id}</div>
                        </div>
                      </td>
                      <td className="px-6 py-4 text-gray-900">
                        {new Date(record.year, record.month - 1).toLocaleDateString('en-US', { year: 'numeric', month: 'long' })}
                      </td>
                      <td className="px-6 py-4 text-gray-900">{formatCurrency(record.basic_salary)}</td>
                      <td className="px-6 py-4 text-green-600">{formatCurrency(record.allowances)}</td>
                      <td className="px-6 py-4 text-red-600">{formatCurrency(record.deductions)}</td>
                      <td className="px-6 py-4 text-gray-900 font-bold">{formatCurrency(record.net_salary)}</td>
                      <td className="px-6 py-4">
                        <span
                          className={`inline-block px-3 py-1 rounded-full text-xs font-medium capitalize ${
                            record.payment_status === 'paid'
                              ? 'bg-green-100 text-green-700 border border-green-200'
                              : record.payment_status === 'pending'
                              ? 'bg-yellow-100 text-yellow-700 border border-yellow-200'
                              : 'bg-red-100 text-red-700 border border-red-200'
                          }`}
                        >
                          {record.payment_status}
                        </span>
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          )}
        </div>
      </div>

      {/* Add Employee Modal */}
      {showAddEmployeeModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-xl font-bold text-gray-900 mb-4">Add New Employee</h2>
            <form onSubmit={handleAddEmployee} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Employee Code *</label>
                <input
                  type="text"
                  required
                  className="input-field"
                  value={employeeForm.employee_code}
                  onChange={(e) => setEmployeeForm({ ...employeeForm, employee_code: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Department *</label>
                <input
                  type="text"
                  required
                  className="input-field"
                  value={employeeForm.department}
                  onChange={(e) => setEmployeeForm({ ...employeeForm, department: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Designation *</label>
                <input
                  type="text"
                  required
                  className="input-field"
                  value={employeeForm.designation}
                  onChange={(e) => setEmployeeForm({ ...employeeForm, designation: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Joining Date *</label>
                <input
                  type="date"
                  required
                  className="input-field"
                  value={employeeForm.joining_date}
                  onChange={(e) => setEmployeeForm({ ...employeeForm, joining_date: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Salary Amount *</label>
                <input
                  type="number"
                  required
                  min="0"
                  step="0.01"
                  className="input-field"
                  value={employeeForm.salary_amount}
                  onChange={(e) => setEmployeeForm({ ...employeeForm, salary_amount: e.target.value })}
                />
              </div>
              <div className="flex space-x-3">
                <button type="submit" className="btn-primary flex-1">
                  Add Employee
                </button>
                <button
                  type="button"
                  onClick={() => setShowAddEmployeeModal(false)}
                  className="btn-secondary flex-1"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Mark Attendance Modal */}
      {showAttendanceModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-xl font-bold text-gray-900 mb-4">Mark Attendance</h2>
            <form onSubmit={handleMarkAttendance} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Employee *</label>
                <select
                  required
                  className="input-field"
                  value={attendanceForm.employee_id}
                  onChange={(e) => setAttendanceForm({ ...attendanceForm, employee_id: e.target.value })}
                >
                  <option value="">Select Employee</option>
                  {employees.map((emp) => (
                    <option key={emp.id} value={emp.id}>
                      {emp.employee_code} - {emp.designation}
                    </option>
                  ))}
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Date *</label>
                <input
                  type="date"
                  required
                  className="input-field"
                  value={attendanceForm.date}
                  onChange={(e) => setAttendanceForm({ ...attendanceForm, date: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Check In Time</label>
                <input
                  type="time"
                  className="input-field"
                  value={attendanceForm.check_in}
                  onChange={(e) => setAttendanceForm({ ...attendanceForm, check_in: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Check Out Time</label>
                <input
                  type="time"
                  className="input-field"
                  value={attendanceForm.check_out}
                  onChange={(e) => setAttendanceForm({ ...attendanceForm, check_out: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Status *</label>
                <select
                  required
                  className="input-field"
                  value={attendanceForm.status}
                  onChange={(e) => setAttendanceForm({ ...attendanceForm, status: e.target.value })}
                >
                  <option value="present">Present</option>
                  <option value="absent">Absent</option>
                  <option value="half_day">Half Day</option>
                  <option value="leave">Leave</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Notes</label>
                <textarea
                  rows="2"
                  className="input-field"
                  value={attendanceForm.notes}
                  onChange={(e) => setAttendanceForm({ ...attendanceForm, notes: e.target.value })}
                />
              </div>
              <div className="flex space-x-3">
                <button type="submit" className="btn-primary flex-1">
                  Mark Attendance
                </button>
                <button
                  type="button"
                  onClick={() => setShowAttendanceModal(false)}
                  className="btn-secondary flex-1"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Generate Payroll Modal */}
      {showPayrollModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-xl font-bold text-gray-900 mb-4">Generate Payroll</h2>
            <form onSubmit={handleGeneratePayroll} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Employee *</label>
                <select
                  required
                  className="input-field"
                  value={payrollForm.employee_id}
                  onChange={(e) => setPayrollForm({ ...payrollForm, employee_id: e.target.value })}
                >
                  <option value="">Select Employee</option>
                  {employees.map((emp) => (
                    <option key={emp.id} value={emp.id}>
                      {emp.employee_code} - {emp.designation}
                    </option>
                  ))}
                </select>
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Month *</label>
                  <select
                    required
                    className="input-field"
                    value={payrollForm.month}
                    onChange={(e) => setPayrollForm({ ...payrollForm, month: parseInt(e.target.value) })}
                  >
                    {Array.from({ length: 12 }, (_, i) => (
                      <option key={i + 1} value={i + 1}>
                        {new Date(2000, i).toLocaleString('en-US', { month: 'long' })}
                      </option>
                    ))}
                  </select>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Year *</label>
                  <input
                    type="number"
                    required
                    min="2000"
                    max="2100"
                    className="input-field"
                    value={payrollForm.year}
                    onChange={(e) => setPayrollForm({ ...payrollForm, year: parseInt(e.target.value) })}
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Allowances</label>
                <input
                  type="number"
                  min="0"
                  step="0.01"
                  className="input-field"
                  value={payrollForm.allowances}
                  onChange={(e) => setPayrollForm({ ...payrollForm, allowances: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Deductions</label>
                <input
                  type="number"
                  min="0"
                  step="0.01"
                  className="input-field"
                  value={payrollForm.deductions}
                  onChange={(e) => setPayrollForm({ ...payrollForm, deductions: e.target.value })}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Notes</label>
                <textarea
                  rows="2"
                  className="input-field"
                  value={payrollForm.notes}
                  onChange={(e) => setPayrollForm({ ...payrollForm, notes: e.target.value })}
                />
              </div>
              <div className="flex space-x-3">
                <button type="submit" className="btn-primary flex-1">
                  Generate Payroll
                </button>
                <button
                  type="button"
                  onClick={() => setShowPayrollModal(false)}
                  className="btn-secondary flex-1"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};

export default Payroll;
