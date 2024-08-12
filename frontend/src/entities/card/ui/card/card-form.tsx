import { CardBody, CardHeader, Heading, Text, Card as CardElement } from '@chakra-ui/react';
import { FC } from 'react';
import { Card } from 'entities/card';

type Props = {
  card: Card;
};

export const CardForm: FC<Props> = ({ card }) => {
  return (
    <CardElement cursor="default" h={160} maxW={400}>
      <CardHeader>
        <Heading size="md">{card.question}</Heading>
      </CardHeader>
      <CardBody>
        <Text fontSize="sm" noOfLines={2}>
          {card.answer}
        </Text>
      </CardBody>
    </CardElement>
  );
};
