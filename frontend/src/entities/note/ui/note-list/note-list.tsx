import { Grid, GridItem } from '@chakra-ui/react';
import { FC } from 'react';
import { INote } from '../../types';
import { Note } from '../note';

interface IProps {
  notes: INote[];
  onNoteClick?: (noteId: INote['id']) => void;
}

export const NoteList: FC<IProps> = ({ notes, onNoteClick }) => {
  const onListNoteClick = (noteId: INote['id']) => () => onNoteClick?.(noteId);

  return (
    <Grid templateColumns="repeat(5, 1fr)" gap={6}>
      {notes.map((note) => (
        <GridItem key={note.id}>
          <Note
            key={note.id}
            note={note}
            onClick={onNoteClick ? onListNoteClick(note.id) : undefined}
          />
        </GridItem>
      ))}
    </Grid>
  );
};
