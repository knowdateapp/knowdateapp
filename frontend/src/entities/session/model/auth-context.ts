import { createContext } from 'react';

export type Workspace = string | null;

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
