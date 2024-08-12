import { Card, CardBody, CardHeader, Heading, Text } from '@chakra-ui/react';
import { FC, MouseEventHandler } from 'react';
import { Note } from 'entities/note/types';

type Props = {
  note: Note;
  onClick?: MouseEventHandler<HTMLDivElement> | undefined;
};

export const NoteForm: FC<Props> = ({ note, onClick }) => {
  return (
    <Card onClick={onClick} cursor={onClick ? 'pointer' : 'default'}>
      <CardHeader>
        <Heading size="md">{note.title}</Heading>
      </CardHeader>
      <CardBody>
        <Text fontSize="sm">Description</Text>
      </CardBody>
    </Card>
  );
};
