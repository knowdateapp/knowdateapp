import { Grid, GridItem } from '@chakra-ui/react';
import { FC } from 'react';
import { INote } from '../../types';
import { Note } from '../note';

// TODO: Абстрактное название. Можно сделать более абстрактно и вынести в общий компонент.
//  По факту такое и для карточек нужно и для заметок.
interface IProps {
  notes: INote[];
  onNoteClick?: (noteId: INote['id']) => void;
}

export const NoteList: FC<IProps> = ({ notes, onNoteClick }) => {
  // TODO: Что это такое?
  const onListNoteClick = (noteId: INote['id']) => () => onNoteClick?.(noteId);

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
