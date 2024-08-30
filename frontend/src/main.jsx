import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import Login from './routes/login/Login.jsx'
import Dashboard from './routes/dashboard/Dashboard.jsx'

import {
  createBrowserRouter, 
  RouterProvider
} from 'react-router-dom'

const router = createBrowserRouter([
  {
    element: <App />,
    children: [
      {
        path: "/login",
        element: <Login />
      },
      {
        path: "/dashboard",
        element: <Dashboard />
      }
    ]
  }
])

createRoot(document.getElementById('root')).render(
    <RouterProvider router={router} />
)
