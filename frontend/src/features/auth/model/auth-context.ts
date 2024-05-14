import { createContext } from 'react';

export interface IAuthContext {
  isAuth: boolean;
  setAuth: (value: boolean) => void;
}
// TODO если контекст понадобится в других фичах перенести его в "entities/session"
export const AuthContext = createContext<IAuthContext>({
  isAuth: false,
  setAuth: () => {
    return;
  },
});
