import { FC } from 'react';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import { Routes } from 'shared/config';
import { useAuth } from '../../model';

export const ProtectedRoute: FC = () => {
  const { isAuth } = useAuth();
  const location = useLocation();
  return isAuth ? <Outlet /> : <Navigate to={Routes.Auth} state={{ from: location }} replace />;
};
