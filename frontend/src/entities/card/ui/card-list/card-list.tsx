import { FC } from 'react';
import { SimpleGrid } from '@chakra-ui/react';
import { CardForm } from 'entities/card';
import { Card } from 'entities/card/types';

type Props = {
  cards: Card[];
};

export const CardList: FC<Props> = ({ cards }) => {
  return (
    <SimpleGrid minChildWidth={200} spacing={6}>
      {cards.map((card) => (
        <CardForm card={card} />
      ))}
    </SimpleGrid>
  );
};
