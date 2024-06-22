import { Box, Heading } from '@chakra-ui/react';
import { FC } from 'react';
import { useParams } from 'react-router-dom';

export const NotePage: FC = () => {
  const { noteId } = useParams();

  return (
    <Box>
      <Heading>Note - {noteId}</Heading>
    </Box>
  );
};
