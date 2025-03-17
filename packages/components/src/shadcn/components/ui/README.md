# UI Components

These components were generated with shadcn.

They can be updated with this command:
`npx shadcn@latest add -o -y <component-names>`.

However, any modifications we made will need to be reapplied.

## Modifications

We've made the following modifications to these shadcn components:

### Alert

- Description color is black instead of gray to make our message stand out.

### Checkbox

- Border is black instead of gray so it is visible on our blue background.

### Dialog

- "X" icon to close the dialog is white to stand out on a black background and
  its size is increased.
- Content padding is changed.

### Form

- Errors messages are red instead of gray, and form labels turn red on error.

### Input

- Background is white for better contrast on our blue background.

### Textarea

- Background is white for better contrast on our blue background.
- Removed default behavior of sizing based on content (`field-sizing-content`)
  so it does not expand to show the entire letter of petitions and push the
  submit button too far down on the page. The user can scroll in the textarea
  to see the entire petition if desired.
