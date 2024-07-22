import { CardBody, CardHeader, Heading, Text, Card as CardElement } from '@chakra-ui/react';
import { FC } from 'react';
import { CardEntity } from 'entities/card';

type Props = {
  card: CardEntity;
};

export const Card: FC<Props> = ({ card }) => {
  return (
    <CardElement cursor="default">
      <CardHeader>
        <Heading size="md">{card.question}</Heading>
      </CardHeader>
      <CardBody>
        {/*TODO: Обработать клик на карточку, по типу перевернуть ее*/}
        <Text fontSize="sm">{card.answer}</Text>
      </CardBody>
    </CardElement>
  );
};
