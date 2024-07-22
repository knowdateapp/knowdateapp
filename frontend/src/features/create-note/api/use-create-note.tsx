import { useMutation } from '@tanstack/react-query';
import { AxiosResponse } from 'axios';
import { CreateNoteParams, NoteEntity, notesController } from 'entities/note';
import { IApiError } from 'shared/api';

export const useCreateNote = () => {
  return useMutation<AxiosResponse<NoteEntity>, IApiError, CreateNoteParams>({
    mutationFn: notesController.createNote,
  });
};
