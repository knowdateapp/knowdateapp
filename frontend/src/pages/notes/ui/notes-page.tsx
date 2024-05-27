import { Heading } from '@chakra-ui/react';
import { FC } from 'react';
import { useAuth } from 'entities/session';

export const NotesPage: FC = () => {
  const { workspace } = useAuth();
  return <Heading>Workspace: {workspace}</Heading>;
};
