export function applyToMatrix<T>(
  matrix: T[][],
  action: (x: number, y: number, value?: T) => void
): void {
  for (const [y, row] of matrix.entries()) {
    for (const [x] of row.entries()) {
      action(x, y);
    }
  }
}
