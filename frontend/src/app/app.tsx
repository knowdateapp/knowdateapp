import { Box, Flex } from '@chakra-ui/react';
import { FC, useState } from 'react';
import { RouterProvider } from 'react-router-dom';
import { Header } from 'widgets/header';
import { AuthContext } from 'features/auth';
import { Providers } from './providers';
import 'shared/config/i18n';
import { router } from './router.tsx';

export const App: FC = () => {
  const [isAuth, setAuth] = useState(false);

  return (
    <Providers>
      <AuthContext.Provider value={{ isAuth, setAuth }}>
        <Flex flexDirection="column" h="100vh" px={4} py={4}>
          <Header />
          <Box flex={1} as="main" pt={8}>
            <RouterProvider router={router} />
          </Box>
        </Flex>
      </AuthContext.Provider>
    </Providers>
  );
};
