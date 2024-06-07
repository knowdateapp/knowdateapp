import { createBrowserRouter, Navigate } from 'react-router-dom';
import { AuthPage } from 'pages/auth';
import { NotePage } from 'pages/note-page';
import { NotesPage } from 'pages/notes';
import { ProtectedRoute } from 'features/auth';
import { Routes } from 'shared/config';

export const router = createBrowserRouter([
  {
    element: <ProtectedRoute />,
    children: [
      {
        path: Routes.Notes,
        element: <NotesPage />,
      },
      {
        path: Routes.Note,
        element: <NotePage />,
      },
    ],
  },
  {
    path: Routes.Auth,
    element: <AuthPage />,
  },
  {
    path: '*',
    element: <Navigate to={Routes.Notes} />,
  },
]);
