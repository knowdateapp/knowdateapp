import { Grid, GridItem } from '@chakra-ui/react';
import { FC } from 'react';
import { Note } from 'entities/note/types';
import { NoteForm } from '../note';

type Props = {
  notes: Note[];
  onNoteClick?: (noteId: Note['id']) => void;
};

export const NoteList: FC<Props> = ({ notes, onNoteClick }) => {
  const onListNoteClick = (noteId: Note['id']) => () => onNoteClick?.(noteId);

  return (
    <Grid templateColumns="repeat(5, 1fr)" gap={6}>
      {notes.map((note) => (
        <GridItem key={note.id}>
          <NoteForm
            key={note.id}
            note={note}
            onClick={onNoteClick ? onListNoteClick(note.id) : undefined}
          />
        </GridItem>
      ))}
    </Grid>
  );
};
