import { Card, CardBody, CardHeader, Heading, Text } from '@chakra-ui/react';
import { FC, MouseEventHandler } from 'react';
import { NoteEntity } from 'entities/note/types';

type Props = {
  note: NoteEntity;
  onClick?: MouseEventHandler<HTMLDivElement> | undefined;
};

export const Note: FC<Props> = ({ note, onClick }) => {
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
