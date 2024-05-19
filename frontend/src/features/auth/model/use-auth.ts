import { useContext } from 'react';
import { AuthContext, IAuthContext } from './auth-context.ts';

export const useAuth = (): IAuthContext => {
  const { isAuth, setAuth } = useContext(AuthContext);

  return { isAuth, setAuth };
};
