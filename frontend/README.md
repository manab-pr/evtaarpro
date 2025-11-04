# EvtaarPro Frontend

Modern React frontend for the EvtaarPro collaboration platform.

## 🚀 Quick Start

### Prerequisites

- Node.js 18+ (or use nvm)
- npm or yarn
- Backend API running on `http://localhost:8080`

### Installation

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The application will open at `http://localhost:3000`

## 📦 Tech Stack

- **React 18** - UI library
- **Vite** - Build tool & dev server
- **React Router v6** - Client-side routing
- **TailwindCSS** - Utility-first CSS
- **Axios** - HTTP client
- **React Hot Toast** - Notifications
- **Lucide React** - Icons
- **date-fns** - Date formatting

## 🎨 Features

### ✅ Authentication
- Login with email/password
- User registration
- JWT token management
- Automatic token refresh
- Protected routes

### ✅ Dashboard
- Welcome screen with user info
- Statistics cards (meetings, users, projects)
- Recent meetings list
- Quick action buttons
- Real-time data

### ✅ User Management
- View current user profile
- Edit profile information
- List all team members
- Search users by name/email
- User cards with role badges
- Pagination support

### ✅ Meetings
- Create new meetings
- List all meetings with pagination
- View meeting details
- Join meetings (opens Jitsi)
- Meeting status indicators
- Date & time display

### ✅ UI/UX
- Responsive design (mobile, tablet, desktop)
- Smooth animations
- Loading states
- Error handling
- Toast notifications
- Beautiful color scheme
- Gradient backgrounds

## 📁 Project Structure

```
frontend/
├── public/              # Static assets
├── src/
│   ├── components/      # Reusable components
│   │   └── Layout.jsx   # Main layout with sidebar
│   ├── context/         # React Context
│   │   └── AuthContext.jsx  # Auth state management
│   ├── pages/           # Page components
│   │   ├── Login.jsx
│   │   ├── Register.jsx
│   │   ├── Dashboard.jsx
│   │   ├── Profile.jsx
│   │   ├── Users.jsx
│   │   ├── Meetings.jsx
│   │   ├── MeetingCreate.jsx
│   │   └── MeetingDetails.jsx
│   ├── services/        # API services
│   │   └── api.js       # Axios instance & API calls
│   ├── App.jsx          # App component with routing
│   ├── main.jsx         # Entry point
│   └── index.css        # Global styles
├── index.html           # HTML template
├── package.json         # Dependencies
├── vite.config.js       # Vite configuration
└── tailwind.config.js   # Tailwind configuration
```

## 🔌 API Integration

The frontend communicates with the backend API through axios:

**Base URL**: `/api/v1` (proxied to `http://localhost:8080`)

**Authentication**: JWT tokens stored in localStorage

**Endpoints**:
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `POST /auth/logout` - Logout user
- `GET /users/me` - Get current user
- `PUT /users/me` - Update profile
- `GET /users` - List users
- `POST /meetings` - Create meeting
- `GET /meetings` - List meetings
- `GET /meetings/:id` - Get meeting details
- `POST /meetings/:id/join` - Join meeting

## 🎨 Customization

### Colors

Edit `tailwind.config.js` to change the color scheme:

```js
theme: {
  extend: {
    colors: {
      primary: {
        // Your custom colors
      },
    },
  },
}
```

### Logo

Replace the text logo in `src/components/Layout.jsx`:

```jsx
<h1 className="text-2xl font-bold text-primary-600">
  YourLogo
</h1>
```

## 📝 Usage Examples

### Login

1. Go to `http://localhost:3000/login`
2. Enter email: `test@example.com`
3. Enter password: `your-password`
4. Click "Sign In"

### Create Meeting

1. Go to Dashboard
2. Click "New Meeting" button
3. Fill in meeting details:
   - Title: "Team Standup"
   - Description: "Daily standup meeting"
   - Start Time: Select date and time
   - Max Participants: 50
4. Click "Create Meeting"

### Join Meeting

1. Go to Meetings page
2. Find your meeting
3. Click on the meeting card
4. Click "Join Meeting"
5. Jitsi will open in a new window

## 🛠️ Development

### Available Scripts

```bash
# Start dev server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

### Environment Variables

Create `.env` file for custom configuration:

```env
VITE_API_URL=http://localhost:8080/api/v1
```

## 🐛 Troubleshooting

### API Connection Issues

**Problem**: Cannot connect to backend
**Solution**:
1. Ensure backend is running on port 8080
2. Check proxy configuration in `vite.config.js`
3. Verify CORS settings in backend

### Token Expiration

**Problem**: Getting 401 errors
**Solution**:
1. Login again to get new token
2. Check token expiration time in backend config

### Build Errors

**Problem**: Build fails
**Solution**:
```bash
# Clean install
rm -rf node_modules package-lock.json
npm install
npm run build
```

## 📱 Responsive Design

The application is fully responsive:

- **Mobile** (< 640px): Single column, hamburger menu
- **Tablet** (640px - 1024px): Two columns
- **Desktop** (> 1024px): Full sidebar, multi-column layouts

## 🎯 Key Features to Test

1. **Authentication Flow**
   - Register new user
   - Login
   - View profile
   - Update profile
   - Logout

2. **Dashboard**
   - View statistics
   - See recent meetings
   - Quick actions

3. **Users**
   - List all users
   - Search functionality
   - View user details

4. **Meetings**
   - Create meeting
   - View meetings list
   - Join meeting (Jitsi opens)
   - View meeting details

## 🚀 Deployment

### Build for Production

```bash
npm run build
```

The `dist/` folder contains the production build.

### Deploy to Netlify

```bash
# Install Netlify CLI
npm install -g netlify-cli

# Deploy
netlify deploy --prod --dir=dist
```

### Deploy to Vercel

```bash
# Install Vercel CLI
npm install -g vercel

# Deploy
vercel --prod
```

## 🎨 UI Screenshots

The app features:
- 🎨 Modern gradient backgrounds
- 🎭 Smooth animations and transitions
- 📊 Beautiful statistics cards
- 🎯 Intuitive navigation
- 📱 Mobile-responsive design
- 🌈 Professional color scheme

## 📄 License

MIT License - Part of EvtaarPro project

## 🤝 Contributing

This is a demonstration project for the EvtaarPro platform.

---

**Enjoy building with EvtaarPro! 🚀**
