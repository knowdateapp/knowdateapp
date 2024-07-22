import { FC } from 'react';
import { Button } from '@chakra-ui/react';
import { useTranslation } from 'react-i18next';
import { generatePath, useNavigate } from 'react-router-dom';
import { useAuth } from 'entities/session';
import { Routes } from 'shared/config';
import { useCreateCard } from 'features/create-card/api';

export const CreateCardButton: FC = () => {
  const navigate = useNavigate();
  const { workspace } = useAuth();
  const { t } = useTranslation();
  const { mutate, isPending } = useCreateCard();

  const onClick = () => {
    mutate(
      { workspace, answer: 'answer', question: 'question' },
      {
        onError: (error) => console.error('Create Card error - ', error),
        onSuccess: (response) => {
          navigate(generatePath(Routes.Card, { cardId: response.data.id }));
        },
      },
    );
  };

  return (
    <Button onClick={onClick} isLoading={isPending}>
      {t('features.createCardButton.createButtonText')}
    </Button>
  );
};
