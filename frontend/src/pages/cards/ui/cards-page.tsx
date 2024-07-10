import { FC } from 'react';
import { useAuth } from '../../../entities/session';
import { useNavigate } from 'react-router-dom';
import { Box } from '@chakra-ui/react';

export const CardsPage: FC = () => {
  const { workspace } = useAuth();
  const navigate = useNavigate();

  workspace;
  navigate;

  return <Box>Cards</Box>;
};
