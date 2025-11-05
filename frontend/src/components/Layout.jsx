import React, { useState } from 'react';
import { Outlet, Link, useLocation, useNavigate } from 'react-router-dom';
import {
  LayoutDashboard,
  Users,
  Video,
  User,
  LogOut,
  Menu,
  X,
  Bell,
  Settings,
  DollarSign,
  Building2
} from 'lucide-react';
import { useAuth } from '../context/AuthContext';

const Layout = () => {
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const { user, logout } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();

  const navigation = [
    { name: 'Dashboard', href: '/dashboard', icon: LayoutDashboard },
    { name: 'Meetings', href: '/meetings', icon: Video },
    { name: 'Users', href: '/users', icon: Users },
    { name: 'Payroll', href: '/payroll', icon: DollarSign },
    { name: 'CRM', href: '/crm', icon: Building2 },
    { name: 'Notifications', href: '/notifications', icon: Bell },
    { name: 'Profile', href: '/profile', icon: User },
  ];

  const handleLogout = async () => {
    await logout();
    navigate('/login');
  };

  const activeNav = navigation.find((item) =>
    location.pathname === item.href || location.pathname.startsWith(`${item.href}/`)
  );

  const SidebarContent = ({ onNavigate }) => (
    <div className="flex h-full flex-col bg-white shadow-xl">
      <div className="flex items-center justify-between border-b border-gray-200 px-6 py-5">
        <h1 className="text-2xl font-bold text-primary-600">EvtaarPro</h1>
        <button
          onClick={onNavigate}
          className="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 lg:hidden"
          aria-label="Close sidebar"
        >
          <X className="h-5 w-5" />
        </button>
      </div>

      <div className="border-b border-gray-200 px-6 py-6">
        <div className="flex items-center space-x-3">
          <div className="flex h-12 w-12 items-center justify-center rounded-full bg-primary-100 text-lg font-semibold text-primary-600">
            {user?.first_name?.[0]}
            {user?.last_name?.[0]}
          </div>
          <div className="min-w-0">
            <p className="truncate text-sm font-semibold text-gray-900">{user?.full_name}</p>
            <p className="truncate text-xs text-gray-500">{user?.email}</p>
          </div>
        </div>
        <span className="mt-3 inline-flex items-center rounded-full bg-primary-50 px-3 py-1 text-xs font-medium capitalize text-primary-700">
          {user?.role}
        </span>
      </div>

      <nav className="flex-1 space-y-1 overflow-y-auto px-4 py-6">
        {navigation.map((item) => {
          const Icon = item.icon;
          const isActive =
            location.pathname === item.href || location.pathname.startsWith(`${item.href}/`);

          return (
            <Link
              key={item.name}
              to={item.href}
              onClick={onNavigate}
              className={`group flex items-center rounded-lg px-4 py-3 text-sm font-medium transition-all ${
                isActive
                  ? 'bg-primary-50 text-primary-700 shadow-sm'
                  : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
              }`}
            >
              <Icon
                className={`mr-3 h-5 w-5 ${
                  isActive ? 'text-primary-600' : 'text-gray-400 group-hover:text-gray-600'
                }`}
              />
              {item.name}
            </Link>
          );
        })}
      </nav>

      <div className="border-t border-gray-200 p-4">
        <button
          onClick={handleLogout}
          className="flex w-full items-center rounded-lg px-4 py-3 text-sm font-medium text-red-600 transition-colors hover:bg-red-50"
        >
          <LogOut className="mr-3 h-5 w-5" />
          Logout
        </button>
      </div>
    </div>
  );

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-white to-slate-100">
      {sidebarOpen && (
        <div className="fixed inset-0 z-40 flex lg:hidden">
          <div
            className="flex-1 bg-gray-900/40 backdrop-blur-sm"
            onClick={() => setSidebarOpen(false)}
          />
          <aside className="relative h-full w-72 animate-slide-in">
            <SidebarContent onNavigate={() => setSidebarOpen(false)} />
          </aside>
        </div>
      )}

      <div className="mx-auto flex min-h-screen w-full max-w-[1400px] px-0 sm:px-4 lg:px-8">
        <aside className="hidden w-72 lg:block lg:py-8">
          <div className="h-full overflow-hidden rounded-3xl border border-gray-100 shadow-lg">
            <SidebarContent onNavigate={() => setSidebarOpen(false)} />
          </div>
        </aside>

        <div className="flex flex-1 flex-col lg:py-8">
          <header className="sticky top-0 z-20 border-b border-gray-200 bg-white/90 backdrop-blur">
            <div className="flex h-16 items-center justify-between px-4 sm:px-6 lg:px-8">
              <div className="flex items-center space-x-3">
                <button
                  onClick={() => setSidebarOpen(true)}
                  className="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 lg:hidden"
                  aria-label="Open sidebar"
                >
                  <Menu className="h-5 w-5" />
                </button>
                <h2 className="text-lg font-semibold text-gray-900">
                  {activeNav?.name || 'Dashboard'}
                </h2>
              </div>

              <div className="flex items-center space-x-3">
                <Link to="/notifications" className="relative rounded-full p-2 text-gray-400 transition-colors hover:bg-gray-100 hover:text-gray-600">
                  <Bell className="h-5 w-5" />
                  <span className="absolute right-2 top-2 h-2 w-2 rounded-full bg-red-500"></span>
                </Link>
                <button className="rounded-full p-2 text-gray-400 transition-colors hover:bg-gray-100 hover:text-gray-600">
                  <Settings className="h-5 w-5" />
                </button>
              </div>
            </div>
          </header>

          <main className="flex-1 px-4 py-6 sm:px-6 lg:px-10">
            <div className="mx-auto max-w-6xl animate-fade-in space-y-6">
              <Outlet />
            </div>
          </main>
        </div>
      </div>
    </div>
  );
};

export default Layout;
