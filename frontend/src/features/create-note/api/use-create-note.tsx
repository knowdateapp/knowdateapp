import { useMutation } from '@tanstack/react-query';
import { AxiosResponse } from 'axios';
import { ICreateNoteParams, INote, notesController } from 'entities/note';
import { IApiError } from 'shared/api';

export const useCreateNote = () => {
  return useMutation<AxiosResponse<INote>, IApiError, ICreateNoteParams>({
    mutationFn: notesController.createNote,
  });
};
