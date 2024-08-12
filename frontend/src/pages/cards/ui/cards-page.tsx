import { FC } from 'react';
import { CardList } from 'entities/card';
import { useAuth } from 'entities/session';
import { useCards } from 'entities/card/api';

export const CardsPage: FC = () => {
  const { workspace } = useAuth();
  const { data } = useCards(workspace);

  return <CardList cards={data ? data.cards : []} />;
};
