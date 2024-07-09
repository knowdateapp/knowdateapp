import { createContext } from 'react';
import { Workspace } from 'shared/model';

export interface IAuthContext {
  isAuth: boolean;
  setWorkspace: (value: Workspace) => void;
  workspace: Workspace;
}

export const AuthContext = createContext<IAuthContext>({
  isAuth: false,
  // TODO: Если тип nullable, то зачем тут использовать отличное от null значение по умолчанию?
  workspace: '',
  setWorkspace: () => {
    return;
  },
});
