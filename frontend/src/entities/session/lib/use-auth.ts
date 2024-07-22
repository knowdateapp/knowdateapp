import { useContext } from 'react';
import { AuthContext, AuthContextType } from '../model';

export const useAuth = (): AuthContextType => {
  return useContext(AuthContext);
};
