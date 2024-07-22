import { Box } from '@chakra-ui/react';
import { FC } from 'react';
import { generatePath, useNavigate } from 'react-router-dom';
import { useNotes, NoteList } from 'entities/note';
import { useAuth } from 'entities/session';
import { Routes } from 'shared/config';

export const NotesPage: FC = () => {
  const { workspace } = useAuth();
  const { data } = useNotes(workspace);
  const navigate = useNavigate();

  const onNoteClick = (noteId: string) => navigate(generatePath(Routes.Note, { noteId }));

  if (!data) {
    return null;
  }

  return (
    <Box>
      <NoteList notes={data.notes} onNoteClick={onNoteClick} />
    </Box>
  );
};
