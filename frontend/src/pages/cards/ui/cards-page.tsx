import { FC } from 'react';
import { Box } from '@chakra-ui/react';
import { CardList } from 'entities/card';
import { useAuth } from 'entities/session';
import { useCards } from 'entities/card/api';

export const CardsPage: FC = () => {
  const { workspace } = useAuth();
  // TODO: Потенциально undefined.
  const { data } = useCards(workspace);

  return (
    <Box>
      <CardList cards={data ? data.cards : []} />
    </Box>
  );
};
