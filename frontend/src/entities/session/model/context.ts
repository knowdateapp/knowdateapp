import { createContext } from 'react';
import { Workspace } from 'shared/model';

export type AuthContextType = {
  isAuth: boolean;
  setWorkspace: (value: Workspace) => void;
  workspace: Workspace;
};

export const AuthContext = createContext<AuthContextType>({
  isAuth: false,
  workspace: '',
  setWorkspace: () => {
    return;
  },
});
