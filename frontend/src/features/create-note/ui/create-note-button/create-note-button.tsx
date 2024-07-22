import { Button } from '@chakra-ui/react';
import { FC } from 'react';
import { useTranslation } from 'react-i18next';
import { generatePath, useNavigate } from 'react-router-dom';
import { useAuth } from 'entities/session';
import { Routes } from 'shared/config';
import { useCreateNote } from 'features/create-note/api';

export const CreateNoteButton: FC = () => {
  const navigate = useNavigate();
  const { t } = useTranslation();
  const { mutate, isPending } = useCreateNote();
  const { workspace } = useAuth();

  const onClick = () => {
    mutate(
      { workspace },
      {
        onError: (error) => console.error('Network error - ', error),
        onSuccess: (response) => {
          navigate(generatePath(Routes.Note, { noteId: response.data.id }));
        },
      },
    );
  };

  return (
    <Button onClick={onClick} isLoading={isPending}>
      {t('features.createNoteButton.createButtonText')}
    </Button>
  );
};
