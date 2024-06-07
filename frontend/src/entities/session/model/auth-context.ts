import { createContext } from 'react';
import { Workspace } from 'shared/model';

export interface IAuthContext {
  isAuth: boolean;
  setWorkspace: (value: Workspace) => void;
  workspace: Workspace;
}

export const AuthContext = createContext<IAuthContext>({
  isAuth: false,
  workspace: '',
  setWorkspace: () => {
    return;
  },
});
