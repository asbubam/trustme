use clap::Command;

fn main() {
    let matches = Command::new("trustme")
        .version("0.1.0")
        .about("a simple CLI to manage Trust Relationship of AWS IAM Role Trust me.\nYou can add and remove your identity to Trust Relastionship easily.")
        .subcommand_required(true)
        .arg_required_else_help(true)
        .subcommand(Command::new("whoami"))
        .get_matches();

    match matches.subcommand() {
        Some(("whoami", _sub_matches)) => println!("whoami",),
        _ => unreachable!("Exhausted list of subcommands and subcommand_required prevents `None`"),
    }
}
