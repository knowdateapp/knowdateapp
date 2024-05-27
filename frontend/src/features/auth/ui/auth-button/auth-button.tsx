import { Button } from '@chakra-ui/react';
import { FC } from 'react';
import { useTranslation } from 'react-i18next';
import { useAuth } from 'entities/session';

export const AuthButton: FC = () => {
  const { t } = useTranslation();
  const { isAuth, setWorkspace } = useAuth();

  const onClick = () => {
    setWorkspace(null);
  };

  return <Button onClick={onClick}>{t(isAuth ? 'logOut' : 'logIn')}</Button>;
};
