import { FC, useState } from 'react';
import { RouterProvider } from 'react-router-dom';
import { AuthContext } from 'entities/session';
import 'shared/config/i18n';
import { Workspace } from 'shared/model';
import { Providers } from './providers';
import { router } from './router.tsx';

export const App: FC = () => {
  const [workspace, setWorkspace] = useState<Workspace>(localStorage.getItem('workspace') || '');

  const onSetWorkSpace = (userWorkspace: Workspace) => {
    setWorkspace(userWorkspace);

    if (userWorkspace) {
      localStorage.setItem('workspace', userWorkspace);
      return;
    }

    localStorage.removeItem('workspace');
  };

  return (
    <Providers>
      <AuthContext.Provider
        value={{ workspace, setWorkspace: onSetWorkSpace, isAuth: Boolean(workspace) }}
      >
        <RouterProvider router={router} />
      </AuthContext.Provider>
    </Providers>
  );
};
