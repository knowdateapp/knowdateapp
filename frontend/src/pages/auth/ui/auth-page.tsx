import { Flex } from '@chakra-ui/react';
import { FC } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthForm, useAuth } from 'features/auth';
import { Routes } from 'shared/config';

export const AuthPage: FC = () => {
  const { isAuth } = useAuth();

  if (isAuth) {
    return <Navigate to={Routes.Main} />;
  }

  return (
    <Flex height="100%" justifyContent="center">
      <AuthForm />
    </Flex>
  );
};
