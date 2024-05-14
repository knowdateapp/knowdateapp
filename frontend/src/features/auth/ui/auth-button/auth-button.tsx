import { Button } from '@chakra-ui/react';
import { FC } from 'react';
import { useTranslation } from 'react-i18next';
import { useAuth } from '../../model';

export const AuthButton: FC = () => {
  const { t } = useTranslation();
  const { isAuth, setAuth } = useAuth();

  const onClick = () => {
    setAuth(!isAuth);
  };

  return <Button onClick={onClick}>{t(isAuth ? 'logOut' : 'logIn')}</Button>;
};
