import { createBrowserRouter, Navigate, Outlet } from 'react-router-dom';
import { AuthPage } from 'pages/auth';
import { NotePage } from 'pages/note-page';
import { NotesPage } from 'pages/notes';
import { Header } from 'widgets/header';
import { ProtectedRoute } from 'features/auth';
import { Routes } from 'shared/config';
import { PageLayout } from 'shared/ui';
import { CardsPage } from '../pages/cards';

export const router = createBrowserRouter([
  {
    element: <PageLayout content={<ProtectedRoute />} header={<Header />} />,
    children: [
      {
        path: Routes.Notes,
        element: <NotesPage />,
      },
      {
        path: Routes.Note,
        element: <NotePage />,
      },
      {
        path: Routes.Cards,
        element: <CardsPage />,
      },
    ],
  },
  {
    element: <PageLayout content={<Outlet />} header={<Header />} />,
    children: [
      {
        path: Routes.Auth,
        element: <AuthPage />,
      },
    ],
  },
  {
    path: '*',
    element: <Navigate to={Routes.Notes} />,
  },
]);
