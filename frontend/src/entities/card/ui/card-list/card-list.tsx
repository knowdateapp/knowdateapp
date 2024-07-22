import { FC } from 'react';
import { Grid, GridItem } from '@chakra-ui/react';
import { Card } from '../card';
import { CardEntity } from 'entities/card/types';

type Props = {
  cards: CardEntity[];
};

export const CardList: FC<Props> = ({ cards }) => {
  return (
    <Grid templateColumns="repeat(5, 1fr)" gap={5}>
      {cards.map((card) => (
        <GridItem key={card.id}>
          <Card key={card.id} card={card} />
        </GridItem>
      ))}
    </Grid>
  );
};
