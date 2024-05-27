import { FC } from 'react';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import { useAuth } from 'entities/session';
import { Routes } from 'shared/config';

export const ProtectedRoute: FC = () => {
  const { isAuth } = useAuth();
  const location = useLocation();
  return isAuth ? <Outlet /> : <Navigate to={Routes.Auth} state={{ from: location }} replace />;
};
