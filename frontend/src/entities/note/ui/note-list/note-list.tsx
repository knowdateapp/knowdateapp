import { Grid, GridItem } from '@chakra-ui/react';
import { FC } from 'react';
import { NoteEntity } from 'entities/note/types';
import { Note } from '../note';

type Props = {
  notes: NoteEntity[];
  onNoteClick?: (noteId: NoteEntity['id']) => void;
};

export const NoteList: FC<Props> = ({ notes, onNoteClick }) => {
  const onListNoteClick = (noteId: NoteEntity['id']) => () => onNoteClick?.(noteId);

  return (
    <Grid templateColumns="repeat(5, 1fr)" gap={6}>
      {notes.map((note) => (
        <GridItem key={note.id}>
          <Note
            key={note.id}
            note={note}
            // TODO: А мы хотим допустить использование этого компонента без onNoteClick?
            onClick={onNoteClick ? onListNoteClick(note.id) : undefined}
          />
        </GridItem>
      ))}
    </Grid>
  );
};
