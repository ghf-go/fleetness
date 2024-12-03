export default function (item, rowIndex) {
  if (rowIndex % 2 == 0) {
    return "warning-row";
  } else {
    return "success-row";
  }
}
