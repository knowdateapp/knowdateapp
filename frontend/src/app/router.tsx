import { createBrowserRouter, Navigate } from 'react-router-dom';
import { AuthPage } from 'pages/auth';
import { MainPage } from 'pages/main';
import { ProtectedRoute } from 'features/auth';
import { Routes } from 'shared/config';

export const router = createBrowserRouter([
  {
    element: <ProtectedRoute />,
    children: [
      {
        path: Routes.Main,
        element: <MainPage />,
      },
    ],
  },
  {
    path: Routes.Auth,
    element: <AuthPage />,
  },
  {
    path: '*',
    element: <Navigate to={Routes.Main} />,
  },
]);
