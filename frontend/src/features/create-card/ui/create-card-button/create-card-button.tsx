import { FC } from 'react';
import { Button } from '@chakra-ui/react';
import { useTranslation } from 'react-i18next';

type Props = {
  onClick: () => void;
};

export const CreateCardButton: FC<Props> = ({ onClick }) => {
  const { t } = useTranslation();

  return <Button onClick={onClick}>{t('features.createCardButton.createButtonText')}</Button>;
};
