import { FC } from 'react';
import { Grid, GridItem } from '@chakra-ui/react';
import { ICard } from '../../types';
import { Card } from '../card';

interface IProps {
  cards: ICard[];
}

export const CardList: FC<IProps> = ({ cards }) => {
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
