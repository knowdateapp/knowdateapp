import {
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
} from '@chakra-ui/react';
import { FC } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
};

export const CreateCardFormModal: FC<Props> = ({ isOpen, onClose }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose} scrollBehavior="outside" size="lg">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Создайте карточку</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <Text>Hello, World!</Text>
        </ModalBody>
        <ModalFooter>
          <Button onClick={onClose} colorScheme="blue">
            Закрыть
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
