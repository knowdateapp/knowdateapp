import { FC } from 'react';
import { Grid, GridItem } from '@chakra-ui/react';
import { Card } from '../card';
import { CardEntity } from 'entities/card/types';

type Props = {
  cards: CardEntity[];
};

export const CardList: FC<Props> = ({ cards }) => {
  return (
    // TODO: Нужно более гибким способом задавать количество элементов в строке.
    <Grid templateColumns="repeat(5, 1fr)" gap={6}>
      {cards.map((card) => (
        <GridItem key={card.id}>
          <Card card={card} />
        </GridItem>
      ))}
    </Grid>
  );
};
