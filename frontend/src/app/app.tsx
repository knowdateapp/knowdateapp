import { Box, Flex } from '@chakra-ui/react';
import { FC, useState } from 'react';
import { RouterProvider } from 'react-router-dom';
import { Header } from 'widgets/header';
import { AuthContext } from 'entities/session';
import { Workspace } from 'shared/model';
import { Providers } from './providers';
import 'shared/config/i18n';
import { router } from './router.tsx';

export const App: FC = () => {
  const [workspace, setWorkspace] = useState<Workspace>(localStorage.getItem('workspace'));

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
        <Flex flexDirection="column" px={4} py={4}>
          <Header />
          <Box flex={1} as="main" pt={8}>
            <RouterProvider router={router} />
          </Box>
        </Flex>
      </AuthContext.Provider>
    </Providers>
  );
};
