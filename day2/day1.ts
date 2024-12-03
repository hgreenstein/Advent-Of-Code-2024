import fs from "fs";

function part1() {
  const reports = fs.readFileSync("./day2.input", "utf-8");
  const splitReports = reports.split("\n");
  const intReports = splitReports.map((report) => {
    const splitReport = report.split(" ");
    return splitReport.map((element) => Number.parseInt(element));
  });
  const validReports = intReports.reduce((acc, currentReport) => {
    // const correctDirection =
    //   currentReport[0] > currentReport[1]
    //     ? (a: number, b: number) => a > b
    //     : (a: number, b: number) => a < b;
    // for (let i = 0; i < currentReport.length - 1; i++) {
    //   if (!correctDirection(currentReport[i], currentReport[i + 1])) {
    //     // console.log(
    //     //   `Direction check failed at index ${i} for report ${currentReport}`
    //     // );
    //     return acc;
    //   }
    //   const absDiff = Math.abs(currentReport[i] - currentReport[i + 1]);
    //   if (absDiff > 3 || absDiff === 0) {
    //     // console.log(
    //     //   `Difference check failed at index ${i} for report ${currentReport}`
    //     // );
    //     return acc;
    //   }
    // }
    // return ++acc;
    if (isValidReport(currentReport)) {
      return ++acc;
    }
    return acc;
  }, 0);
  console.log(`Valid reports for part 1 are ${validReports}\n`);
}

function part2() {
  const reports = fs.readFileSync("./day2.input", "utf-8");
  const splitReports = reports.split("\n");
  const intReports = splitReports.map((report) => {
    const splitReport = report.split(" ");
    return splitReport.map((element) => Number.parseInt(element));
  });
  const validReports = intReports.reduce((acc, currentReport) => {
    for (let index = 0; index < currentReport.length; index++) {
      if (
        isValidReport(
          [...currentReport].filter((_, innerIndex) => innerIndex !== index)
        )
      ) {
        return ++acc;
      }
    }
    return acc;
  }, 0);
  console.log(`Valid reports for part 2 are ${validReports}\n`);
}
part1();
part2();

function isValidReport(report: number[]): boolean {
  const correctDirection =
    report[0] > report[1]
      ? (a: number, b: number) => a > b
      : (a: number, b: number) => a < b;
  for (let i = 0; i < report.length - 1; i++) {
    const absDiff = Math.abs(report[i] - report[i + 1]);
    if (
      !correctDirection(report[i], report[i + 1]) ||
      absDiff > 3 ||
      absDiff === 0
    ) {
      // console.log(
      //   `Direction check failed at index ${i} for report ${report}`
      // );
      return false;
    }
  }
  return true;
}
