import { useContext } from 'react';
import { AuthContext, IAuthContext } from '../model';

export const useAuth = (): IAuthContext => {
  const { isAuth, setWorkspace, workspace } = useContext(AuthContext);

  return { isAuth, setWorkspace, workspace };
};
